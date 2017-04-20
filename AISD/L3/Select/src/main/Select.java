package main;

import java.util.*;

/**
 * Created by Bartosz on 20.04.2017
 */
public class Select {

    private static final int ARRAY_SIZE = 32;
    private static final int NUMBER_BOUNDS = 10000;
    private static final int NUMBER_OF_TESTS = 10;
    private static int comparisonR = 0;
    private static int comparisonS = 0;

    private static int[] generateRandomArray(int size, int bound) {
        int[] ret = new int[size];
        Random r = new Random();
        for(int i = 0; i<size; i++) {
            ret[i] = r.nextInt(bound);
        }
        return ret;
    }

    private static int[] generateDiffRandomArray(int size){
        ArrayList<Integer> list = new ArrayList<>();
        int[] ret = new int[size];
        for(int i = 1; i <= size; i++) {
            list.add(i);
        }
        Collections.shuffle(list);
        for(int i = 0; i<size; i++){
            ret[i] = list.get(i);
        }

        return ret;
    }

    private static void printArray(int[] array) {
        for(int i: array) {
            System.out.print(i+" ");
        }
        System.out.print("\n");
    }

    private static int randomized_Select(int[] array, int start, int end, int i) throws Exception {
        if(start == end) {
            return array[start];
        }
        if(i==0) return -1;
        if(start < end) {

            int q = randomized_partition(array, start, end+1);
            int k = q - start + 1;
            if (i == k) return array[q];
            else if (i < k) return randomized_Select(array, start, q - 1, i);
            else return randomized_Select(array, q + 1, end, i - k);
        }
        return -1;
    }

    private static int randomized_partition(int[] A, int p, int q) throws Exception {
        if(p==q) return p;
        if(p>q) throw new Exception("Cos jest nie tak");
        Random r = new Random();
        int pivotIdx = p + r.nextInt((q-p));
        swap(A,p,pivotIdx);
        int pivot  = A[p];
        // i to indeks pivota
        int i = p, tmp;
        for(int j=p+1; j<q; j++) {
            comparisonR++;
            if(A[j] <= pivot) {

                //swap A[j],a[i+1]
         //       quickSwapKeys++;
                i++;
                tmp = A[j];
                A[j] = A[i];
                A[i] = tmp;
            }
        }
        //quickSwapKeys++;
        tmp = A[p];
        A[p] = A[i];
        A[i] = tmp;

        return i;
    }

    private static void swap(int[] array, int i, int j) {
        int tmp = array[j];
        array[j] = array[i];
        array[i] = tmp;
    }

    private static int select(int[] array, int i, int p, int q) {
//        System.out.println("Rozpoczynam select na p "+p+" q "+ q);
       int pivot = findMedian(array,p,q);
        if((q-p) < 6) {
            return array[i-1];
        }
       int idx = partition(array, pivot,p, q);
       if((idx+1) == i) return array[idx];
       else if(i < idx+1) {
           return select(array,i,p,idx);
       } else return select(array,i,idx+1,q);
    }

    private static int partition(int[] A, int value, int p, int q) {
//        System.out.println("Robie partition");
        int index = 0;
        for(int i = p; i < q; i++) {
            if(A[i]==value) {
                index = i;
                break;
            }
        }
        swap(A,p,index);
        int pivot  = A[p];
        // i to indeks pivota
        int i = p, tmp;
        for(int j=p+1; j<q; j++) {
                comparisonS++;
            if(A[j] <= pivot) {

                //swap A[j],a[i+1]
                //       quickSwapKeys++;
                i++;
                tmp = A[j];
                A[j] = A[i];
                A[i] = tmp;
            }
        }
        //quickSwapKeys++;
        tmp = A[p];
        A[p] = A[i];
        A[i] = tmp;
        return i;
    }

    private static int findMedian(int array[], int p, int q){
        if(q-p == 1) return array[p];
        if((q-p) < 6) {
            sort(array, p, q);
            return array[((q-p)-1)/2];
        } else {
            //sortujemy piątkami
            for(int j = p; j < q; j += 5) {
                sort(array,j,j+5);
            }
            int median_size = (q-p)/5;
            int size_of_over = 0;
            if((q-p)%5 != 0) {
                median_size+=1;
                size_of_over = (q-p)%5;
            }
            int[] medians = new int[median_size];
            for(int j = 0; j<medians.length && (j*5+2)<(q-p)-size_of_over; j++){
                medians[j] = array[j*5+p+2];
            }
            if(size_of_over != 0) {
                int idx = q - (size_of_over)/2-1;
                medians[medians.length-1] = array[idx];
            }
            return findMedian(medians,0,medians.length);
        }
    }

    private static void sort(int[] array, int i, int j) {
        int koniec = j;

        if(koniec >= array.length) {
            koniec = array.length;
        }
        if(i-j == 0) {
            return;

        }
        for (int k = i+1; k < koniec; k++) {
            int key = array[k];
            int l = k-1;
            while(l >= i && key < array[l]) {
                comparisonS++;
                array[l+1]=array[l];
                l--;
            }
            comparisonS++;
            array[l+1] = key;
        }
    }

    private static void test() throws Exception {
        Random r = new Random();
        for(int i = 1; i<NUMBER_OF_TESTS; i+=100){
            int min = 2147483647;
            int max = 0;
            int avg = 0;
            for(int j =0; j<100; j++){
                int[] A = generateDiffRandomArray(i);
                comparisonR = 0;
                randomized_Select(A,0,A.length-1, r.nextInt(A.length)+1);
                avg += comparisonR;
                if(comparisonR > max) max = comparisonR;
                if(comparisonR < min) min = comparisonR;
            }
            avg /= 100;
            System.out.print(i+" "+min+" "+avg+" "+max+ " ");

            min = 2147483647;
            max = 0;
            avg = 0;
            for(int j = 0; j<100; j++){
                int[] A = generateDiffRandomArray(i);
                comparisonS = 0;
                select(A, r.nextInt(A.length)+1,0, A.length);
                avg += comparisonS;
                if(comparisonS > max) max = comparisonS;
                if(comparisonS < min) min = comparisonS;
            }
            avg /= 100;
            System.out.print(min+" "+avg+" "+max);
            System.out.print("\n");
        }
    }

    public static void main(String[] args) throws Exception {
        test();

        System.out.println("1.Losowy ciag");
        System.out.println("2.Losowy ciag roznowartosciowy");
        Scanner input = new Scanner(System.in);
        int n = input.nextInt();
        switch (n){
            case 1:
                System.out.print("Podaj rozmiar danych ");
                int v = input.nextInt();
                System.out.print("Podaj numer szukanej statystyki pozycyjnej");
                int k = input.nextInt();
                los(v, k);
                break;
            case 2:
                System.out.print("Podaj rozmiar danych ");
                v = input.nextInt();
                System.out.print("Podaj numer szukanej statystyki pozycyjnej");
                k = input.nextInt();
                roz(v,k);
        }
    }

    private static void roz(int v, int k) throws Exception {
        int[] tab = generateDiffRandomArray(v);
        int[] copy = new int[v];
        for(int i = 0; i < copy.length; i++) {
            copy[i] = tab[i];
        }

        System.out.println("Randomized select: "+ randomized_Select(tab,0,tab.length-1,k));
        System.out.println("Select: "+ select(copy,k,0,copy.length) );
        Arrays.sort(tab);
        printArray(tab);
    }

    private static void los(int v, int k) throws Exception {
        int[] tab = generateRandomArray(v, 10000);
        int[] copy = new int[v];
        for(int i = 0; i < copy.length; i++) {
            copy[i] = tab[i];
        }

        System.out.println("Randomized select: "+ randomized_Select(tab,0,tab.length-1,k));
        System.out.println("Select: "+ select(copy,k,0,copy.length) );
        Arrays.sort(tab);
        printArray(tab);

    }
}