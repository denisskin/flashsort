flashsort
=========

The algorithm is based on the method of bucket sort (bin sort).
The algorithm has a run time close to linear and uses O(1) memory.

Example
-------
``` c
// sort array of strings
char *names[10] = {
    "Hunter",
    "Isaac",
    "Christopher",
    "Bob",
    "Faith",
    "Alice",
    "Gabriel",
    "Denis",
    "****",
    "Ethan",
};

static const char* get_char(const void *value, unsigned pos) {
    return *((char*)value + pos)? (char*)value + pos : NULL;
}

flashsort((void**)names, 10, get_char);

```

``` c
// sort integer values
int nums[10] = {9, 6, 7, 0, 3, 1, 3, 2, 5, 8};
flashsort_const(nums, 10, sizeof(int), sizeof(int));

```

``` c
// sort key-value array by key
typedef struct {
    int  key;
    char *value;
} KeyValue;

KeyValue names[10] = {
    {8, "Hunter"},
    {9, "Isaac"},
    {3, "Christopher"},
    {2, "Bob"},
    {6, "Faith"},
    {1, "Alice"},
    {7, "Gabriel"},
    {4, "Denis"},
    {0, "none"},
    {5, "Ethan"},
};

flashsort_const(names, 10, sizeof(KeyValue), sizeof(names->key));

```

Define sort-function by C-macro

``` c
// define function flashsort_int by macro
#define FLASH_SORT_NAME  flashsort_int
#define FLASH_SORT_TYPE  int
#include "../src/flashsort_macro.h"

int nums[10] = { 9, 6, 7, 0, 3, 1, 3, 2, 5, 8 };
flashsort_int(nums, 10);

```
 

Benchmarks
----------
Benchmarks of sorting of strings.

run
``` shell
cd benchmarks
gcc ../src/*.c benchmarks.c  -o benchmarks.o && ./benchmarks.o
```
output
``` txt

Benchmarks sorting of strings
 
--------------------------------------------------------------------------------------------
Count		Flash-sort		          	     Quick-sort
elements	total time	   one operation    total time	 one operation
--------------------------------------------------------------------------------------------
    147		0.000011 sec	[0.726 mcsec]	0.000009 sec	[0.628 mcsec]	-13.58%
    274		0.000025 sec	[0.921 mcsec]	0.000025 sec	[0.923 mcsec]	+0.30 %
    512		0.000065 sec	[1.272 mcsec]	0.000067 sec	[1.317 mcsec]	+3.53 %
    955		0.000114 sec	[1.197 mcsec]	0.000144 sec	[1.503 mcsec]	+25.64%
   1782		0.000207 sec	[1.160 mcsec]	0.000326 sec	[1.831 mcsec]	+57.83%
   3326		0.000371 sec	[1.116 mcsec]	0.000618 sec	[1.859 mcsec]	+66.67%
   6208		0.000779 sec	[1.255 mcsec]	0.001446 sec	[2.329 mcsec]	+85.58%
  11585		0.001705 sec	[1.472 mcsec]	0.002720 sec	[2.348 mcsec]	+59.52%
  21618		0.003395 sec	[1.571 mcsec]	0.005210 sec	[2.410 mcsec]	+53.43%
  40342		0.007023 sec	[1.741 mcsec]	0.010266 sec	[2.545 mcsec]	+46.17%
  75281		0.012347 sec	[1.640 mcsec]	0.020921 sec	[2.779 mcsec]	+69.44%
 140479		0.025313 sec	[1.802 mcsec]	0.045267 sec	[3.222 mcsec]	+78.83%


Benchmarks sorting of integers

-------------------------------------------------------------------------------------------
Count		Flash-sort		          	    Quick-sort
elements	total time	   one operation	total time	 one operation
-------------------------------------------------------------------------------------------
    100		0.000013 sec	[1.318 mcsec]	0.000004 sec	[0.440 mcsec]	-66.60%
    194		0.000016 sec	[0.839 mcsec]	0.000009 sec	[0.460 mcsec]	-45.16%
    374		0.000026 sec	[0.701 mcsec]	0.000027 sec	[0.713 mcsec]	+1.62 %
    724		0.000046 sec	[0.636 mcsec]	0.000065 sec	[0.904 mcsec]	+42.21%
   1398		0.000094 sec	[0.674 mcsec]	0.000149 sec	[1.065 mcsec]	+58.08%
   2702		0.000232 sec	[0.859 mcsec]	0.000361 sec	[1.335 mcsec]	+55.38%
   5220		0.000423 sec	[0.810 mcsec]	0.000664 sec	[1.271 mcsec]	+57.02%
  10085		0.000922 sec	[0.914 mcsec]	0.001420 sec	[1.408 mcsec]	+54.05%
  19483		0.002306 sec	[1.183 mcsec]	0.002970 sec	[1.524 mcsec]	+28.82%
  37640		0.004492 sec	[1.193 mcsec]	0.006220 sec	[1.653 mcsec]	+38.46%
  72716		0.007396 sec	[1.017 mcsec]	0.012951 sec	[1.781 mcsec]	+75.11%
 140479		0.012348 sec	[0.879 mcsec]	0.026624 sec	[1.895 mcsec]	+115.62%
 271388		0.020366 sec	[0.750 mcsec]	0.052127 sec	[1.921 mcsec]	+155.95%
 524288		0.039443 sec	[0.752 mcsec]	0.104651 sec	[1.996 mcsec]	+165.32%
 

Benchmarks sorting of IP-addresses log

-------------------------------------------------------------------------------------------------------
Count		Flash-sort			Quick-sort
elements	total time	one operation	total time	one operation
-------------------------------------------------------------------------------------------------------
    107		0.000007 sec	[0.685 mcsec]	0.000009 sec	[0.832 mcsec]	+21.50%
    208		0.000020 sec	[0.968 mcsec]	0.000016 sec	[0.790 mcsec]	-18.39%
    407		0.000027 sec	[0.675 mcsec]	0.000041 sec	[1.017 mcsec]	+50.59%
    793		0.000056 sec	[0.712 mcsec]	0.000111 sec	[1.402 mcsec]	+96.99%
   1547		0.000093 sec	[0.604 mcsec]	0.000241 sec	[1.558 mcsec]	+158.07%
   3016		0.000214 sec	[0.709 mcsec]	0.000704 sec	[2.335 mcsec]	+229.21%
   5881		0.000443 sec	[0.754 mcsec]	0.001472 sec	[2.502 mcsec]	+232.08%
  11466		0.000768 sec	[0.670 mcsec]	0.002829 sec	[2.467 mcsec]	+268.31%
  22354		0.001343 sec	[0.601 mcsec]	0.005789 sec	[2.590 mcsec]	+330.90%
  43584		0.002602 sec	[0.597 mcsec]	0.010738 sec	[2.464 mcsec]	+312.67%
  84974		0.005628 sec	[0.662 mcsec]	0.022314 sec	[2.626 mcsec]	+296.50%
 165670		0.010292 sec	[0.621 mcsec]	0.045589 sec	[2.752 mcsec]	+342.97%
 323000		0.018893 sec	[0.585 mcsec]	0.096238 sec	[2.980 mcsec]	+409.39%
 629739		0.037370 sec	[0.593 mcsec]	0.196754 sec	[3.124 mcsec]	+426.51%
 
```