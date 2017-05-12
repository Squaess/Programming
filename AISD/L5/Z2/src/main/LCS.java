package main;

import java.util.Random;

/**
 * Created by Bartosz on 11.05.2017
 */

public class LCS {

    public static void main (String[] args) {

//        int[] A = new int[]{1,2,3,2,4,1,2};
//        int[] B = new int[]{2,4,3,1,2,1};
//
//        Lcs_Structure structure = new Lcs_Structure(A,B);
//        System.out.println(structure.getLength());
//        structure.print_LCS(A.length-1, B.length-1);

        Random r = new Random();
        int[] X;
        int[] Y;
        for (int n = 1; n < 1000; n++) {

            int cmp = 0;
            int mem = 0;

            for (int c = 0; c < 100; c++) {

                int divide = r.nextInt(n);
                X = new int[n - divide];
                Y = new int[divide];
                int comparison = 0;
                int memory = 0;
                for (int i = 0; i < 1000; i++) {

                    for (int j = 0; j < X.length; j++) {
                        X[j] = r.nextInt();
                    }

                    for (int j = 0; j < Y.length; j++) {
                        Y[j] = r.nextInt();
                    }


                    Lcs_Structure s = new Lcs_Structure(X, Y);

                    comparison += s.comparison;
                    memory += s.m * s.n * 2;
                }

                comparison /= 1000;
                memory /= 1000;
                cmp += comparison;
                mem += memory;
            }
            cmp /= 100;
            mem /= 100;

            System.out.print(n + " " + cmp + " " + mem + "\n");
        }
    }

}
