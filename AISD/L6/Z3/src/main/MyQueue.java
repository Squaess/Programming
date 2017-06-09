package main;

/**
 * Created by Bartosz on 24.05.2017
 */

public class MyQueue {
    public Vertex[] A;
    private int heapSize;

    public MyQueue(Vertex[] Arr){
        this.heapSize = Arr.length;
        A = Arr.clone();
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
        if(l < heapSize && A[l].compareTo(A[i]) == -1){
            smallest = l;
        } else smallest = i;
        if (r < heapSize && A[r].compareTo(A[smallest]) == -1) {
            smallest = r;
        }

        if (smallest != i) {
            swap(i,smallest);
            heapify(smallest);
        }
    }

    private void swap(int a, int b) {
        Vertex v1 = A[a];
        A[a] = A[b];
        A[b] = v1;
    }

    void print(){
        for (int i= 0; i<A.length; i++) System.out.print("("+A[i].id+" "+A[i].weight+") ");
        System.out.print("\n");
    }

    public Vertex heap_max(){
        return A[0];
    }

    public Vertex heap_Extraxt_Max() throws Exception {
        if(heapSize<1) {
            throw new Exception("Pusta kolejka");
        }
        Vertex ret = A[0];
        A[0] = A[heapSize-1];
        heapSize--;
        heapify(0);
        return ret;
    }

    public void insert(Vertex value) throws Exception {
        if(heapSize >= A.length){
            throw new Exception("Queue is full");
        }
        heapSize++;
        int i = heapSize;
        while( i>0 && A[parent(i)].compareTo(value) == 1) {
            A[i] = A[parent(i)];
            i = parent(i);
        }
        A[i] = value;
    }

    private void decreaseKey(int index, double newKey){
        A[index].weight = newKey;
        System.out.println("\t\tDecrease key z wartoscia "+newKey+" na pozycji "+index);
        print();
        int p = parent(index);
        int smallest;

        if(p < heapSize && A[p].compareTo(A[index]) == -1){
            smallest = p;
        } else smallest = index;

//        if (smallest == index) {
//            swap(index,p);
//         //   decreaseKey(p,newKey);
//        }
        while(A[p].compareTo(A[index]) == 1){
            swap(index,p);
            index = p;
            p =parent(p);
        }
    }

    public void update(int id, double newWeight){

        for(int i = 0; i < heapSize; i++){
            if(A[i].id == id){
                decreaseKey(i, newWeight);
                break;
            }
        }
    }

}
