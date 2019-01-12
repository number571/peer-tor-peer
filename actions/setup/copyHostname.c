#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

#include "../../settings/values.h"

extern void copy_hostname(void) {
	#ifdef DEBUG
		printf("[system_call]: 'cp %s %s'\n", PRIVATE_DIR "hostname", "../");
	#endif
	system("cp " PRIVATE_DIR "hostname ../");
}
