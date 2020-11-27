package phymem

/*
#include <stdio.h>
#include <sys/param.h>
#include <sys/sysctl.h>
#include <sys/types.h>
#include <unistd.h>

static int64_t
getRSS(void)
{
    int ret;
    size_t len;

    struct kinfo_proc2 kp;
    int mib[] = {CTL_KERN, KERN_PROC2, KERN_PROC_PID, getpid(), sizeof(kp), 1};
    size_t miblen = __arraycount(mib);
    len = sizeof(kp);
    ret = sysctl(mib, miblen, &kp, &len, NULL, 0);
    if (ret < 0) {
        return -1;
    }

    long pagesize = sysconf(_SC_PAGESIZE);

    return (int64_t)(kp.p_vm_rssize * pagesize);
}
*/
import "C"
import "fmt"

// for test.
const providedCurrent = true

func Current() (uint, error) {
	rss := int(C.getRSS())
	if rss < 0 {
		return 0, fmt.Errorf("failed to get RSS via sysctl: %d", rss)
	}
	return uint(rss), nil
}
