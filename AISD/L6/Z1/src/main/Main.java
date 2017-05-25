package main;

/**
 * Created by Bartosz on 25.05.2017
 */
public class Main {

    public static void main(String[] args) throws Exception {
        MyQueue heap = new MyQueue(new int[]{50,22,69,87,15,2,31,45,12,6,-4,5,8,7,4,32,54,1,1,50});
        heap.print();
        heap.decreaseKey(1,90);
        heap.print();
    }
}
