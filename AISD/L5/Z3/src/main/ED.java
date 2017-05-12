package main;

import java.util.Scanner;

/**
 * Created by Bartosz on 11.05.2017
 */
public class ED {

    private static char[][] dictionary = new char[][]{
            {'P','A','R','E','N','T'},
            {'P','O','P'},
            {'P','A','I','R'},
            {'P','A','C','K'},
            {'P','A','D'},
            {'P','A','I','N'},
            {'P','A','R','K'},
            {'P','A','P','E','R'},
            {'P','O','L','Y','N','O','M','I','A','L'}
    };

    public static void main (String[] args) {
        System.out.println("Hello");
        char[] c1 = new char[]{'P','O','L','Y','N','O','M','I','A','L'};
        char[] c2 = new char[]{'E','X','P','O','N','E','N','T','I','A','L'};
        char[] c3 = new char[]{'S','N','O','W','Y'};
        char[] c4 = new char[]{'S','U','N','N','Y'};
        ED_Structure s = new ED_Structure(c1,c2);
        System.out.println(s.getDistance());
        ED_Structure s2 = new ED_Structure(c3, c4);
        System.out.println(s2.getDistance());

        Scanner input = new Scanner(System.in);
        String curr_word = input.next();
        char[] curr_char_arr = curr_word.toCharArray();
        int[] closest_distance = new int[3];
        String[] a = new String[]{"","",""};
        closest_distance[2] = Integer.MAX_VALUE;
        closest_distance[1] = Integer.MAX_VALUE;
        closest_distance[0] = Integer.MAX_VALUE;

        for(char[] c: dictionary) {
            ED_Structure ed = new ED_Structure(curr_char_arr, c);
            if(ed.getDistance() < closest_distance[2]) {
                if (ed.getDistance() < closest_distance[1]) {
                    if(ed.getDistance() < closest_distance[0]) {
                        closest_distance[2] = closest_distance[1];
                        closest_distance[1] = closest_distance[0];
                        closest_distance[0] = ed.getDistance();
                        a[2] = a[1];
                        a[1] = a[0];
                        a[0] = String.valueOf(c);

                    } else {
                        closest_distance[2] = closest_distance[1];
                        closest_distance[1] = ed.getDistance();
                        a[2] = a[1];
                        a[1] = String.valueOf(c);
                    }
                } else {
                    closest_distance[2] = ed.getDistance();
                    a[2] = String.valueOf(c);
                }
            }
        }
        for (String s1 : a ) System.out.println(s1);
    }
}
