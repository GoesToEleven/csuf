#include <stdint.h>

typedef struct {
  int64_t* ptr;
  int64_t len;
  int64_t cap;
  int64_t* ret;
} slice;

void sum(slice slc) {
  int64_t total = 0;
  for (int64_t i = 0; i < slc.len; i++) {
    total += slc.ptr[i];
  }
  *slc.ret = slc.len;
}
