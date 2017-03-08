/* Bartosz Banasik Static FIFO */

#include<stdio.h>

struct Node {
	struct Node *next;
	struct Node *prev;
	int val;
};

struct Node first;

int main() {
    first.val=5;
}
