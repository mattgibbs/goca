#include "cago.h"
#include "_cgo_export.h"
void handleConnection_cgo(struct connection_handler_args args) {
	handleConnection(args.chid, args.op);
}

void handleEvent_cgo(struct event_handler_args args) {
	handleEvent(args.usr, args.chid, args.type, args.count, args.dbr, args.status);
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

int ca_create_subscription_cgo(chtype type, unsigned long count, chid chan, long mask, void* callback, void* userarg, evid* pevid) {
	return ca_create_subscription(type, count, chan, mask, callback, userarg, pevid);
}