#include <string.h>	/* memcpy */
#include <pthread.h>            /* __thread */
#include "_cgo_export.h"

static __thread int tls;

void
setTLS(int v)
{
	tls = v;
}

int
getTLS()
{
	return tls;
}

void
echo()
{
    Ohce();
}
