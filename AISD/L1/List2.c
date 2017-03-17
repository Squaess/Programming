#include <stdlib.h>
#include <time.h>
#include <stdio.h>

int size = 0;
typedef struct node {
    int value;
    struct node *next;
    struct node *prev;
} node_t;

void add(node_t *head, int val) {
    
    node_t *new = malloc(sizeof(node_t));
    new->value = val;

    if(head->next == NULL){
        head->next = new;
        new->prev = head;
        new->next=head;
        head->prev = new;
    } else {
        head->next->prev = new;
        new->next = head->next;
        head->next = new;
        new->prev=head;
    }

    size++;
}

void delete(node_t **head, int n) {
    if(n>size) return;

    int c = size/2;                 //sprawdzamy w ktora strone bedzie szybciej
    
    node_t * curr  = (*head);
    node_t * tmp1 = NULL;
    node_t * tmp2 = NULL;


    if(n>c) {                       //idziemy w next
        for(int i = size; i >= n; i--) {
            curr = curr->next;
        }
    } else {
        for(int i = 0; i < n; i++) {
            curr = curr->prev;
        }
    }

    if(curr->next != *head && curr->prev != *head){
        curr->next->prev = curr->prev;
        curr->prev->next = curr->next;
        free(curr);
        size--; 
    } else {
        if(curr->next == *head) {
            curr->prev->next = curr->next;
           (*head)->prev = curr->prev;
            free(curr);
            size--;
        } else {
            curr->next->prev = *head;
            (*head)->next = curr->next;
            free(curr);
            size--;
        }
    }
}

void lprintf(node_t *head) {
    node_t *tmp = head;

    while(tmp->prev != head){
        tmp = tmp->prev;
        printf("%d ",tmp->value);
        if(tmp == NULL){
            break;
        }
       
    }
   
}

int get(node_t * head, int n){

     node_t *tmp = head;

     int c = size/2;
     if(n>c) {
        for(int j = size-n; j>=0; j--) {

            tmp = tmp->next;
        }
        return tmp->value;
     } else {
        for(int j = 0; j<n; j++) {
            tmp = tmp->prev;
        }
        return tmp->value;
     }

}

void merge(node_t * head, node_t * sec) {
        sec->next->prev = head;
        sec->prev->next = head->next;
        head->next->prev = sec->prev;
        head->next = sec;
       free(sec);
}

int main() {
    
    node_t *head = malloc(sizeof(node_t));
    node_t *sec = malloc(sizeof(node_t));
    srand(time(NULL));
    int r = 0;
    for(int i = 0; i<2000; i++) {
        r = (rand());
        add(head, r);
    }

    for(int i = 0; i<2000; i++) {


        float second = 0.0;
        clock_t start;
        clock_t end;
        for (int j = 0; j < 1000; j++) {
            start = clock();
            int n = get(head, i);
            end = clock();
            second += (float)(end-start);
        }
        second = second/1000;
        printf("%d %.5f ",i,second);
        second = 0.0;

        for(int j = 0; j<1000; j++) {

                r = (rand())%2000;
                start = clock();
                int n = get(head,r);
                end = clock();
                second += (float)(end-start);

        }
        second = second/1000;
        printf(" %.5f\n ",second);
    }

    add(head,12);
    add(head,14);
    add(sec,22);
    add(sec,23);

    merge(head, sec);
//    lprintf(head);

}
