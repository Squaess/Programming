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

    int parent(int i){
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
        int largest;
        if(l < heapSize && A[l] > A[i]){
            largest = l;
        } else largest = i;
        if (r < heapSize && A[r] > A[largest]) {
            largest = r;
        }
        if (largest != i) {
            swap(i,largest);
            heapify(largest);
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
}
