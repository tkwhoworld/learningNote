#include <stdio.h>

int sum (int *a ,int n);

int array[2] = {1,2};
int main(){
  printf("begin run program!\n");
  int val = sum(array,2);

  printf("the val result is:%d\n",val);
  return val;
}

