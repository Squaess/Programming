package main;

import static java.lang.Integer.min;

/**
 * Created by aedd on 5/11/17
 */

class ED_Structure {
    private int[][] e;
    private int m, n;
    private char[] X;
    private char[] Y;

    ED_Structure(char[] X, char[] Y) {
        this.m = X.length;
        this.n = Y.length;
        this.X = X;
        this.Y = Y;
        this.e = new int[m+1][n+1];
        compute();
    }

    private void compute() {

        for (int i = 0; i <= m; i++) {
            e[i][0] = i;
        }

        for (int i = 0; i <= n; i++) {
            e[0][i] = i;
        }

        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                int i1 = e[i][j+1]+1;
                int i2 = e[i+1][j]+1;
                int i3 = e[i][j] + diff(i,j);
                e[i+1][j+1] =  min( min(i1, i2), i3);
            }
        }
    }

    private int diff(int i, int j) {
        if (X[i] == Y[j]) return 0;
        else return 1;
    }

    int getDistance() {
        return e[m][n];
    }

    void printFirst() {
        for (int i = 0; i <= m; i++) {
            System.out.print(e[i][0]+" ");
        }
        System.out.print("\n");
        for (int i = 0; i <= n; i++) {
            System.out.println(e[0][i]);
        }
    }
}
