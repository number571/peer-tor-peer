#include <stdio.h>
#include <unistd.h>

#include "headers/work.h"
#include "../../settings/values.h"

extern void initial_setup(void) {
	const char * const dirs[3] = {
		MAIN_DIR,
		CHOICE_DIR,
		SETTINGS_DIR,
	};

	#ifdef DEBUG
		printf("[chdir]: '%s'\n", "/");
	#endif
	chdir("/");

	for (char i = 0; i < 3; ++i) {
		if (!dir_is_exist(dirs[i]))
			create_dir(dirs[i]);

		#ifdef DEBUG
			printf("[chdir]: '%s'\n", dirs[i]);
		#endif
		chdir(dirs[i]);
	}

	const char * const files[2] = {
		CONFIG_FILE,
		BLACK_LIST,
	};

	for (char i = 0; i < 2; ++i) {
		if (!file_is_exist(files[i]))
			create_file(files[i]);
	}
}
