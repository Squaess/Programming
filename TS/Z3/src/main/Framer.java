package main;

import java.io.*;
import java.util.ArrayList;
import java.util.Scanner;

/**
 * Created by aedd on 5/10/17
 */
public class Framer {

    private static final String readFile_Name = "./Z.txt";
    private static final String writeFile_Name = "./W.txt";
    private static final String FLAG = "01111110";

    private static String checkSum_CRC8(String input) {
        StringBuffer frame = new StringBuffer();
        frame.append(input);

        String divider = "11000011";

        frame.append("00000000");

        for (int i = 0; i < frame.length() - 8; i++) {
            String s = frame.substring(i,i+8);
            String o = "";
            if(s.startsWith("1")) {

                for (int j = 0; j < divider.length(); j++) {
                    o += s.charAt(j)^divider.charAt(j) ;
                }

                frame.replace(i,i+8, o);
            }
        }
        return frame.toString().substring(frame.length()-8);
    }

    public static void main(String[] args) {

        Scanner in = new Scanner(System.in);
        boolean keep_going = true;

        while(keep_going) {
            System.out.print("\033[H\033[2J");
            System.out.println("1. Zakoduj");
            System.out.println("2. Odkoduj");

            int choice = in.nextInt();

            switch (choice) {
                case 1:
                    String file = readFile(readFile_Name);
                    code_and_write(file);
                    keep_going = false;
                    break;
                case 2:
                    file = readFile(writeFile_Name);
                    decode_and_write(file);
                    keep_going = false;
                    break;
                default:
                    System.out.println("Zly wybor");
                    try {
                        Thread.sleep(2000);
                    } catch (InterruptedException e) {
                        e.printStackTrace();
                    }
            }
        }

    }

    private static void decode_and_write(String data) {

        BufferedWriter out = null;
        try {
            out = new BufferedWriter(new FileWriter("./Z_copy.txt"));
        } catch (IOException e) {
            e.printStackTrace();
        }
        String to_write;
        // Usuwamy flagi
        data = data.replace(FLAG+FLAG, " ");
        data = data.replace( FLAG,"");

        String[] frames = data.split(" ");

        for(int i = 0; i < frames.length; i++) {
            frames[i] = frames[i].replace("111110","11111");
        }
        try {
            for (String s : frames) {
                String crc = s.substring(s.length() - 8);
                String frame = s.substring(0, s.length() - 8);
                if (crc.equals(checkSum_CRC8(frame))) {
                    to_write = frame;
                } else {
                    to_write = "Ramka " + frame + "posiada nieprawidłowy crc";
                }
                out.write(to_write + "\n");
            }
            out.close();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    private static String biteStuffing(String data) {
        ArrayList<Character> chars = new ArrayList<Character>();
        char[] c = data.toCharArray();
        for (char ch: c){
            chars.add(ch);
        }
        int number_of_ones = 0;
        for(int i = 0; i < chars.size(); i++){
            if (chars.get(i) == '0') number_of_ones =0;
            else number_of_ones++;
            if(number_of_ones == 5) chars.add(i+1, '0');
        }
        String ret = "";
        while(!chars.isEmpty()) {
            ret = chars.get(chars.size()-1) + ret;
            chars.remove(chars.size()-1);
        }
        return ret;
    }

    private static void code_and_write(String data) {

        //otwieramy strumień do zapisu
        BufferedWriter out = null;
        try {
            out = new BufferedWriter(new FileWriter("./W.txt"));
        } catch (IOException e) {
            e.printStackTrace();
        }

        String frame;
        int m = 32;
        try {
            //Dzielimy strumien bitów na porcję po 32 bity
            for (int n = 0; n < data.length(); n += m) {
                //jesli nie ma 32 to bierzemy to co zostało
                if (n + m > data.length()) {
                    m = data.length() - n;
                }
                frame = data.substring(n, n + m);
                /*
                 * Obliczamy pole kontrolne CRC
                 */
                String crc = checkSum_CRC8(frame);
                frame += crc;
                /*
                 * Stosujemy metody bite stuffing
                 */
                frame = biteStuffing(frame);
                /*
                 * Dodajemy flagi poczatku i konca
                 */
                frame = "01111110" + frame + "01111110";

                out.write(frame);
                out.write("\n");

            }
            out.close();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    private static String readFile(String readFile_name) {
        String ret = "";
        String tmp;

        File inputFile = new File(readFile_name);

        try {
            BufferedReader reader = new BufferedReader( new FileReader(inputFile));
            while( (tmp = reader.readLine() ) != null) {
                ret += tmp;
            }

            reader.close();
        } catch (FileNotFoundException e) {
            System.out.println("File not found");
            e.printStackTrace();
            System.exit(0);
        } catch (IOException e) {
            e.printStackTrace();
        }

        if(ret.isEmpty()) {
            System.out.println("Input file is empty");
            System.exit(0);
        }
        return ret;
    }
}
