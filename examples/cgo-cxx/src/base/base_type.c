// int float double long

#include "base_type.h"

API_EXPORT(int) base_add(int a, int b)
{
    return a + b;
}

API_EXPORT(float) base_multi(float a, float b)
{
    return a * b;
}

API_EXPORT(double) base_sub(double a, double b)
{
    return a - b;
}

API_EXPORT(long) base_div(long a, long b)
{
    return (b==0) ? (-1) : (a/b); 
}

API_EXPORT(int) base_sum(int* ptr, int count) 
{
	int total = 0;
	for (int idx = 0; idx < count; idx++) {
		total += ptr[idx];
	}
	return total;
}

// 未导出的方法，在js中使用时，会报错 base_no_export is not a function
int base_no_export(int a, int b)
{
    return (b==0) ? (-1) : (a/b); 
}