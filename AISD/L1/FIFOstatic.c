/* Bartosz Banasik Static FIFO */

#include <stdio.h>

#define MAX 100000

int fifo[MAX], first = 0, last = 0;

int add(int t) {
    if(last>=MAX) {
        return 0;
    }

    fifo[last]=t;
    last = last + 1;

    return 1;
}

int get() {
    if(first<last) {
        int n = first;
        first = first + 1;
        return fifo[n];
    } else{
        printf("No more items in queue");
    }
}

int main() {
    int n;
    while(n!=123) {
        scanf("%d",&n);
        add(n); 
    }
    while (first < last) {
        printf("Kolejny argument: %d\n", get());
    }
}
