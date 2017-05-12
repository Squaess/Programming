package main;

/**
 * Created by Bartosz on 11.05.2017
 */

class Lcs_Structure {
    private int[] X;
    private int[][] c;
    private Direction[][] b;
    int m, n;
    int comparison;

    Lcs_Structure(int[] X, int[] Y) {
        comparison = 0;
        this.X = X;
        this.m = X.length;
        this.n = Y.length;
        this.c = new int[m+1][n+1];
        this.b = new Direction[m][n];
        length(X, Y);
    }

    private void length(int[] X, int[] Y){

        for (int i = 0; i <= m; i++) {
            c[i][0] = 0;
        }

        for (int i = 0; i <= n ; i++) {
            c[0][i] = 0;
        }

        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if(X[i]==Y[j]){
                    comparison ++;
                    c[i+1][j+1] = c[i][j]+1;
                    //jakies przypisanie do tablicy b ktore bedzie pokazywac jak wrocic
                    b[i][j] = Direction.ARROW_CORNER;
                } else if (c[i][j+1] >= c[i+1][j]) {
                    comparison++;
                    c[i+1][j+1] = c[i][j+1];
                    b[i][j] = Direction.ARROW_UP;
                    //do b w górę
                } else {
                    comparison ++;
                    c[i+1][j+1] = c[i+1][j];
                    b[i][j] = Direction.ARROW_LEFT;
                    // do b ze w lewo
                }
            }
        }
    }

    int getLength(){
        return c[m][n];
    }

    void print_LCS(int i, int j) {
        if (i < 0 || j < 0) return;
        if(b[i][j] == Direction.ARROW_CORNER) {
            print_LCS(i-1,j-1);
            System.out.print(X[i]);
        } else if (b[i][j] == Direction.ARROW_UP) print_LCS(i-1, j);
        else print_LCS(i, j-1);
    }
}
