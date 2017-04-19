package main;

import java.util.Arrays;
import java.util.Random;

import static java.lang.StrictMath.log;

/**
 * Created by Aedd on 4/17/17
 */
public class Radix_Sort {
    private static int numberOfSwap = 0;
    private static final int AVG_TEST = 10;
    private static final int N = 100000;
    private static final int STEP = 100;
    private static final int RANGE_OF_RANDOM_DATA = 100;

    private static int getMax(int arr[], int n) {
        if(n==0) return 0;
        int mx = arr[0];
        for (int i = 1; i < n; i++)
            if (arr[i] > mx)
                mx = arr[i];
        return mx;
    }

    private static int[] generateRandomArray(int size) {
        int[] ret = new int[size];
        Random r = new Random();
        for(int i = 0; i < size; i++) {
            ret[i] = r.nextInt(RANGE_OF_RANDOM_DATA);
        }
        return ret;
    }

    private static void printArray(int[] array) {
        for(int i : array) {
            System.out.print(i+" ");
        }
        System.out.print("\n");
    }

    private static void radix_Sort(int[] array) {

        int exp =1;

        int max = getMax(array,array.length);
        for(; max/exp > 0; exp *= 10){
            sort(array,exp, 10);
        }

    }

    private static void sort(int[] array, int exp, int sizeOfData) {
        int[] output = new int[array.length];
        int[] count = new int[sizeOfData];
        Arrays.fill(count, 0);
        int size = array.length;

        for (int anArray : array) {
            count[(anArray / exp) % sizeOfData]++;
        }

        for(int i = 1; i < count.length; i++) {
            count[i] += count[i-1];
        }

        for(int i = size-1; i>=0; i--){
            output[count[(array[i]/exp)%sizeOfData]-1] = array[i];
            numberOfSwap++;
            count[(array[i]/exp)%sizeOfData]--;
        }

        for(int i = 0; i < size; i++) {
            array[i] = output[i];
        }
    }

    private static void betterRadix(int[] array, int b){
        int exp =1;
        int base = b;

        int max = getMax(array,array.length);
        for(; max/exp > 0; exp *= base){
            sort(array,exp, base);
        }
    }

    private static void test() {

        for(int i = 0; i < N; i += STEP) {

            int base = (int)Math.round(log(i)/log(2));
            base = 2^base;

            long time = 0;
            int swaps = 0;
            for(int j = 0; j < AVG_TEST; j++) {
                int[] array = generateRandomArray(i);
                numberOfSwap = 0;
                long startTime = System.nanoTime();
                radix_Sort(array);
                long estimatedTime = System.nanoTime() - startTime;
                time += estimatedTime;
                swaps+= numberOfSwap;
            }
            swaps /= AVG_TEST;
            time /= AVG_TEST;
            System.out.print(i+": "+time+" "+swaps+" ");


            time = 0;
            swaps = 0;
            for(int j = 0; j < AVG_TEST; j++) {

                int[] array = generateRandomArray(i);
                numberOfSwap = 0;
                long startTime = System.nanoTime();
                betterRadix(array, base);
                long estimatedTime = System.nanoTime() - startTime;
                time += estimatedTime;
                swaps+=numberOfSwap;
            }
            swaps /= AVG_TEST;
            time /= AVG_TEST;
            System.out.print(time+" "+swaps+" "+base+"\n");
        }
    }

    public static void main(String[] args) {
       test();
    }
}
