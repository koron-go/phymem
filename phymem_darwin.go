package phymem

// #include <mach/mach.h>
//
//static vm_size_t getRSS(void) {
//    struct task_basic_info info;
//    mach_msg_type_number_t size = sizeof(info);
//    kern_return_t kerr = task_info(mach_task_self(), TASK_BASIC_INFO, (task_info_t)&info, &size);
//    return kerr == KERN_SUCCESS ? info.resident_size : 0;
//}
import "C"
import "errors"

// for test.
const providedCurrent = true

func Current() (uint, error) {
	rss := uint(C.getRSS())
	if rss == 0 {
		return 0, errors.New("failed to get RSS")
	}
	return rss, nil
}
