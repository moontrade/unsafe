#ifndef __MOONTRADE_WORKER_H
#define __MOONTRADE_WORKER_H

#include <assert.h>
#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <pthread.h>

#ifdef __cplusplus
extern "C" {
#endif

#include <stdio.h>
#include <time.h>
#include <unistd.h>

void moontrade_stub();

typedef void moontrade_trampoline_handler(size_t arg0, size_t arg1);

void moontrade_trampoline(size_t fn, size_t arg0, size_t arg1);

void moontrade_sleep(size_t arg0, size_t arg1);

#ifdef __cplusplus
} // extern "C"
#endif

#endif // #ifdef __MOONTRADE_WORKER_H
