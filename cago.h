#include <cadef.h>

void handleConnection_cgo(struct connection_handler_args args);
void handleEvent_cgo(struct event_handler_args args);

//The following are proxy functions that use void pointers instead of types defined in cadef.h
int ca_create_channel_cgo(const char* pvname, void* callback, void* puser, int priority, chid* pchid);
int ca_context_create_cgo(int preemptive);
int ca_create_subscription_cgo(chtype type, unsigned long count, chid chan, long mask, void* callback, void* userarg, evid* pevid);