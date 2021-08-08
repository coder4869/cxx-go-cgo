// int float double long bool

#include "array_type.h"

API_EXPORT(int) array_sum_int(int* ptr, int count) 
{
	int total = 0;
	for (int idx = 0; idx < count; idx++) {
		total += ptr[idx];
	}
	return total;
}


API_EXPORT(float) array_sum_float(float* ptr, int count)
{
	float total = 0;
	for (int idx = 0; idx < count; idx++) {
		total += ptr[idx];
	}
	return total;
}