%{
#include <stdio.h>
#include <gsl/gsl_statistics.h>
%}
double gsl_stats_mean(const double data[], size_t stride, size_t n);
