/* Banasik Bartosz dynamic Fifo Queue */

#include <stdio.h>

typedef struct Node {                               // Struktura wezla wartosc i wskazanie na nastepny element kolejki
    int value;                                      
    struct Node *next;
}Node;

Node *head = NULL;                                  // poczatek kolejki
Node *tail = NULL;                                  // ostatni element kolejki


void add(int n){                                    // dodawanie elementu do kolejki
    if(head==NULL) {                                
        head = (Node *)malloc(sizeof(Node));        // jesli nie mamy jeszcze zadnego elementu w kolejce tworzymy jej pierwszy
        head->value = n;                            // i jednoczesnie ostatni element
        tail = head; 
    } else {                                        // w przeciwnym razie tworzymy nowy wezel o podanej wartosci
                                                    
        Node *new;
        new = (Node *)malloc(sizeof(Node *));
        new->value = n;                             
        new->next = NULL;                           //
        tail->next=new;
        tail=new; 
       
    }
}

void get() {
    Node *temp = head;
    if(head==NULL) {
        printf("Brak elementow\n");
    } else if(head->next==NULL) {
        printf("Brak elementow\n");
        head==NULL;   
    } else {
        head = head->next;
        free(temp);
    }
}

void show() {
    Node *tmp = head;
    if(head==NULL) {
        printf("Brak elemento(show)\n");
    } else while(tmp!=NULL) {
        printf("%d;  ", tmp->value);
        tmp = tmp->next;
    }
    printf("\n");
}

int main() {

    add(12);
    add(13);
    add(120);
    add(12387);
    add(777);
    get();
    get();
    get();
    show();
}
