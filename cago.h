#include <cadef.h>

void handleConnection_cgo(struct connection_handler_args args);
int ca_create_channel_cgo(const char* pvname, void* callback, void* puser, int priority, chid* pchid);
int ca_context_create_cgo(int preemptive);