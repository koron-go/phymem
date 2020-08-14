package phymem

// #include <stdio.h>
// #include <unistd.h>
// #include <sys/types.h>
// #include <sys/sysctl.h>
// #include <sys/param.h>
// #include <sys/user.h>
//
//static segsz_t getRSS(void) {
//    int mib[4];
//    size_t len;
//    struct kinfo_proc ki;
//
//    len = sizeof(mib) / sizeof(mib[0]);
//    if (sysctlnametomib("kern.proc.pid", mib, &len) == -1) {
//        return -1;
//    }
//    mib[3] = getpid();
//    len = sizeof(ki);
//    if (sysctl(mib, sizeof(mib) / sizeof(mib[0]), &ki, &len, NULL, 0) == -1) {
//        return -2;
//    }
//    return ki.ki_rssize;
//}
import "C"
import "fmt"

// for test.
const providedCurrent = true

func Current() (uint, error) {
	rss := int(C.getRSS())
	if rss < 0 {
		return 0, fmt.Errorf("failed to get RSS via sysctl: %d", rss)
	}
	return uint(rss) * 4096, nil
}
