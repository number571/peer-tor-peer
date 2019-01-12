#include <stdio.h>
#include <stdlib.h>
#include <dirent.h>

#include "../../../settings/values.h"
#include "../headers/bool.h"

extern bool dir_is_exist(const char * const name) {
	DIR * const dir = opendir(name);

	if (dir != NULL) {
		closedir(dir);

		#ifdef DEBUG
			printf("[dir_is_exist]: '%s'\n", name);
		#endif
		return true;
	}

	#ifdef DEBUG
		printf("[dir_is_not_exist]: '%s'\n", name);
	#endif
	return false;
}

extern void create_dir(const char * const dir) {
	char buffer[BUFF_SIZE];
	sprintf(buffer, "mkdir %s", dir);

	#ifdef DEBUG
		printf("[system_call]: '%s'\n", buffer);
	#endif
	system(buffer);
}
