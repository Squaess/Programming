package main;

import graph.Edge;

/**
 * Created by Bartosz on 24.05.2017
 */

public class MyQueue {
    private Edge[] A;
    private int heapSize;

    public MyQueue(Edge[] Arr){
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
        if(l < heapSize && A[l].weight < A[i].weight){
            smallest = l;
        } else smallest = i;
        if (r < heapSize && A[r].weight < A[smallest].weight) {
            smallest = r;
        }
        if (smallest != i) {
            swap(i,smallest);
            heapify(smallest);
        }
    }

    private void swap(int a, int b) {
        Edge v1 = A[a];
        A[a] = A[b];
        A[b] = v1;
    }

    void print(){
        for (Edge i : A) System.out.print(i.src+"---"+i.dest+" "+i.weight+"\n");
    }

    public Edge heap_max(){
        return A[0];
    }

    public Edge heap_Extraxt_Max() throws Exception {
        if(heapSize<1) {
            throw new Exception("Pusta kolejka");
        }
        Edge ret = A[0];
        A[0] = A[heapSize-1];
        heapSize--;
        heapify(0);
        return ret;
    }

    public void insert(Edge value) throws Exception {
        if(heapSize>= A.length){
            throw new Exception("Queue is full");
        }
        heapSize++;
        int i = heapSize;
        while(i>0 && A[parent(i)].weight > value.weight){
            A[i] = A[parent(i)];
            i = parent(i);
        }
        A[i] = value;
    }

    public void decreaseKey(int index, int newKey){
        A[index].weight = newKey;
        heapify(index);
    }
}
