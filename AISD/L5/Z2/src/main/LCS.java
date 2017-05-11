package main;

/**
 * Created by Bartosz on 11.05.2017
 */

public class LCS {

    private static int LCS_Length(int[] X, int[] Y) {
        int m = X.length;
        int n = Y.length;
        int[][] c = new int[m+1][n+1];
        Direction[][] b = new Direction[m][n];

        for (int i = 0; i < m; i++) {
            c[i][0] = 0;
        }

        for (int i = 0; i < n ; i++) {
            c[0][i] = 0;
        }
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if(X[i]==Y[j]){
                    c[i+1][j+1] = c[i][j]+1;
                    //jakies przypisanie do tablicy b ktore bedzie pokazywac jak wrocic
                    b[i][j] = Direction.ARROW_CORNER;
                } else if (c[i][j+1] >= c[i+1][j]) {
                    c[i+1][j+1] = c[i][j+1];
                    b[i][j] = Direction.ARROW_UP;
                    //do b w górę
                } else {
                    c[i+1][j+1] = c[i+1][j];
                    b[i][j] = Direction.ARROW_LEFT;
                    // do b ze w lewo
                }
            }
        }
        return c[m][n];
    }

    private void print_LCS(Direction[][] b, int[] X, int i, int j) {
        if (i == 0 || j == 0) return;
        if(b[i][j] == Direction.ARROW_CORNER) {
            print_LCS(b,X,i-1,j-1);
            System.out.print(X[i]);
        } else if (b[i][j] == Direction.ARROW_UP) print_LCS(b, X, i-1, j);
        else print_LCS(b, X, i, j-1);
    }

    public static void main (String[] args) {
        System.out.println("kocham Anie");
        System.out.println("PRawda");
        int[] X = new int[]{1,2,3,2,4,1,2};
        int[] Y = new int[]{2,4,3,1,2,1};

        Lcs_Structure structure = new Lcs_Structure(X,Y);
        System.out.println(structure.getLength());
        structure.print_LCS(X.length-1, Y.length-1);
    }
}
