#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <pthread.h>
#include <stdio.h>
#include <time.h>
#include <unistd.h>
#include <thread>
#include <chrono>
#include "trampoline.h"

void moontrade_stub() {}

//typedef void moontrade_trampoline_handler(size_t arg0, size_t arg1);

void moontrade_trampoline(size_t fn, size_t arg0, size_t arg1) {
	((moontrade_trampoline_handler*)fn)(arg0, arg1);
}

void moontrade_sleep(size_t arg0, size_t arg1) {
	std::this_thread::sleep_for((std::chrono::nanoseconds)arg0);
}