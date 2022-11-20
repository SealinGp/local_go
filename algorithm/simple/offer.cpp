#include <stdio.h>

int NumberOf1(int n);

int main()
{

  if (1)
  {
    printf("result:%d \n", NumberOf1(11));
    return 0;
  }

  printf("1\n");

  return 0;
}

int NumberOf1(int n)
{
  int count = 0;
  unsigned int flag = 1;
  while (flag)
  {
    if (n & flag)
      count++;
    flag = flag << 1;
  }

  return count;
}