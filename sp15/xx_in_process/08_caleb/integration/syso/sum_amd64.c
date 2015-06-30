struct slice {
  void* ptr;
  int64_t len;
  int64_t cap;
}

int64_t SumV2(slice slc) {
  return slc.len;
}
