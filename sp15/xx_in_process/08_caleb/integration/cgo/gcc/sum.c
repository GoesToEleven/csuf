long long sum(long long *xs, long long sz) {
  long long total = 0;
  for (long long i = 0; i < sz; i++) {
    total += xs[i];
  }
  return total;
}
