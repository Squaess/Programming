package main;

/**
 * Created by Bartosz on 24.05.2017
 */

class MyQueue {
    private int[] A;
    private int heapSize;

    MyQueue(int[] Arr){
        this.heapSize = Arr.length;
        this.A = Arr;
        for (int i = (Arr.length/2)-1; i>= 0; i--){
            heapify(i);
        }
    }

    private int parent(int i){
        return (i-1)/2;
    }

    private int left(int i){
        return 2*i+1;
    }

    private int right(int i){
        return 2*i+2;
    }

    private void heapify(int i){
        int l = left(i);
        int r = right(i);
        int smallest;
        if(l < heapSize && A[l] < A[i]){
            smallest = l;
        } else smallest = i;
        if (r < heapSize && A[r] < A[smallest]) {
            smallest = r;
        }
        if (smallest != i) {
            swap(i,smallest);
            heapify(smallest);
        }
    }

    private void swap(int a, int b) {
        int v1 = A[a];
        A[a] = A[b];
        A[b] = v1;
    }

    void print(){
        for (int i : A) System.out.print(i+" ");
        System.out.print("\n");
    }

    int heap_max(){
        return A[0];
    }

    int heap_Extraxt_Max() throws Exception {
        if(heapSize<1) {
            throw new Exception("Pusta kolejka");
        }
        int ret = A[0];
        A[0] = A[heapSize-1];
        heapSize--;
        heapify(0);
        return ret;
    }

    void insert(int value) throws Exception {
        if(heapSize>= A.length){
            throw new Exception("Queue is full");
        }
        heapSize++;
        int i = heapSize;
        while(i>0 && A[parent(i)]>value){
            A[i] = A[parent(i)];
            i = parent(i);
        }
        A[i] = value;
    }

    void decreaseKey(int index, int newKey){
        A[index] = newKey;
        heapify(index);
    }
}
