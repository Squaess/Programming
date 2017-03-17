#include <stdio.h>
#include <time.h>
#include <stdlib.h>

typedef struct node {
    int value;
    struct node *next;
} node_t;

node_t *root = NULL;

void push(node_t *head, int val) {
    
    node_t *tmp = head;
    
    while(tmp->next != NULL) {
        tmp = tmp->next;
    }

    tmp->next = malloc(sizeof(node_t));
    tmp->next->value = val;
    tmp->next->next = NULL;
}

void merge(node_t *first, node_t *sec){
    
    node_t *tmp = first->next;
    
    while(tmp->next != NULL) {
        tmp = tmp->next;
    }
    tmp->next = sec->next;
    return;
}

void lprintf(node_t *head) {
    node_t *tmp = head->next;
    if(tmp->value == NULL){
        printf("Brak elementow");
        free(tmp);
        return;
    }

    while(1){
        printf("%d ",tmp->value);
        tmp = tmp->next;
        if(tmp == NULL){
            break;
        }
       
    }
    free(tmp);
}

int get(int n, node_t * head) {
    node_t * curr = head;
    for(int i = 0; i<n; i++){
        curr = curr->next;
    }
    return curr->value;
}

int main() {
    node_t *head = malloc(sizeof(node_t));
    head->value = NULL;
    node_t *sec = malloc(sizeof(node_t));
    srand(time(NULL));
    int r = 0;
    for(int i = 0; i<2000; i++) {
        r = (rand());
        push(head, r);
    }
    
    for(int i = 0; i<2000; i++) {
        
         
        float second = 0.0;
        clock_t start;
        clock_t end;
        for (int j = 0; j < 1000; j++) {
            start = clock();
            int n = get(i,head);
            end = clock();
            second += (float)(end-start);
        }
        second = second/1000;
        printf("%d %.5f ",i,second);
        second = 0.0;
        
        for(int j = 0; j<1000; j++) {
        
                r = (rand())%2000;
                start = clock();
                int n = get(r,head);
                end = clock();
                second += (float)(end-start);
                
        }
        second = second/1000;
        printf(" %.5f\n ",second);
    }
    push(head, 2);
    push(head, 2);
    push(head, 2);
    push(head, 2);
    push(head, 2);
    push(head, 2);

    push(sec, 3);
    push(sec, 3);
    push(sec, 2);

    merge(head,sec);


    return 0;
}
