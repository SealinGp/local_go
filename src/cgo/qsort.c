#include <stdio.h>
#include <stdlib.h>

//计算数组元素的个数
#define DIM(x) (sizeof(x)/sizeof(x[0]))

static int cmp (const void* a,const void* b) {
      const int* pa = (int*)a;
      const int* pb = (int*)b;
      return *pa-*pb;
}
int main () {
     int values[] = {42,8,109,97,23,25};
     int i;
     qsort(values,DIM(values),sizeof(values[0]),cmp);

     printf("len:%lu, size:%lu \r\n",DIM(values),sizeof(values[0]));

     for (i = 0 ; i < DIM(values); i++) {
           printf("%d\r\n",values[i]);
     }
     return 0;
}
