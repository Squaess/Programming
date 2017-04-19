package main;

import java.util.Arrays;
import java.util.Random;

/**
 * Created by Aedd on 4/17/17
 */
public class Radix_Sort {

    private static final int N = 1000000;
    private static final int SIZE_OF_RANDOM_DATA = 100000;

    static int getMax(int arr[], int n) {
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
            ret[i] = r.nextInt(SIZE_OF_RANDOM_DATA);
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
            sort(array,exp);
        }

    }

    private static void sort(int[] array, int exp) {
        int[] output = new int[array.length];
        int[] count = new int[10];
        Arrays.fill(count, 0);
        int size = array.length;

        for (int anArray : array) {
            count[(anArray / exp) % 10]++;
        }

        for(int i = 1; i < count.length; i++) {
            count[i] += count[i-1];
        }

        for(int i = size-1; i>=0; i--){
            output[count[(array[i]/exp)%10]-1] = array[i];
            count[(array[i]/exp)%10]--;
        }

        for(int i = 0; i < size; i++) {
            array[i] = output[i];
        }
    }

    public static void main(String[] args) {
        System.out.println("Test");
        int[] array = generateRandomArray(N);
        printArray(array);
        radix_Sort(array);
        System.out.println("----------------------------");
        printArray(array);
    }
}
