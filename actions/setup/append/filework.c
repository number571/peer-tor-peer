#include <stdio.h>

#include "../../../settings/values.h"
#include "../headers/bool.h"

extern bool file_is_exist(const char * const name) {
	FILE * const file = fopen(name, "r");
	if (file != NULL) {
		fclose(file);

		#ifdef DEBUG
			printf("[file_is_exist]: '%s'\n", name);
		#endif
		return true;
	}

	#ifdef DEBUG
		printf("[file_is_not_exist]: '%s'\n", name);
	#endif
	return false;
}

extern void create_file(const char * const name) {
	FILE * const file = fopen(name, "w");
	if (file != NULL) {
		#ifdef DEBUG
			printf("[file_created]: '%s'\n", name);
		#endif
		fclose(file);
	}
}
