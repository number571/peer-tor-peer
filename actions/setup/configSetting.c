#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <unistd.h>

#include "../../settings/values.h"
#include "headers/bool.h"

static bool string_is_exist(const char * const name, const char * const str) {
    FILE * const file = fopen(name, "r");
    bool result = false;

    if (file != NULL) {
        char buffer[BUFF_SIZE];
        while (fgets(buffer, BUFF_SIZE, file) != NULL)
            if (!strcmp(buffer, str)) {
                result = true;
                break;
            }
        fclose(file);
    }

    #ifdef DEBUG
        switch (result) {
            case true:  printf("[string_is_exist]: '%s' = %s", name, str); 
                        break;
            case false: printf("[string_is_not_exist]: '%s' = %s", name, str);
        }
    #endif

    return result;
}

static void append_file(const char * const name, const char * const data) {
    FILE * const file = fopen(name, "a");

    if (file != NULL) {
        #ifdef DEBUG
            printf("[append_in_file]: '%s' = %s", name, data);
        #endif

        fputs(data, file);
        fclose(file);
    }
}

extern void config_setting(void) {
    const char * const strings[2] = {
        HIDDEN_SERVICE_DIR  "\n",
        HIDDEN_SERVICE_PORT "\n",
    };

    for (char i = 0; i < 2; ++i)
        if (!string_is_exist(TORRC_FILE, strings[i]))
            append_file(TORRC_FILE, strings[i]);

    #ifdef DEBUG
        printf( "[system_call]: '%s'\n"\
                "[system_call]: '%s'\n",
                "systemctl start tor.service", 
                "systemctl restart tor.service"
        );
    #endif

    system("systemctl start tor.service");
    system("systemctl restart tor.service");

    sleep(1);
}
