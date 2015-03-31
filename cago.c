#include "cago.h"
#include "_cgo_export.h"
void handleConnection_cgo(struct connection_handler_args args) {
	handleConnection(args.chid, args.op);
}

int ca_create_channel_cgo(const char* pvname, void* callback, void* puser, int priority, chid* pchid) {
	return ca_create_channel(pvname, callback, puser, priority, pchid);
}

int ca_context_create_cgo(int preemptive) {
	if (preemptive == 1) {
		return ca_context_create(ca_enable_preemptive_callback);
	}
	
	return ca_context_create(ca_disable_preemptive_callback);
}