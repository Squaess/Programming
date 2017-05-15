package main;

import java.io.*;
import java.nio.charset.Charset;
import java.util.ArrayList;
import java.util.zip.CRC32;
import java.util.zip.Checksum;

/**
 * Created by aedd on 5/10/17
 */
public class Framer {

    private static final Charset UTF8Charset = Charset.forName( "UTF-8" );

    /**
     * Calc CRC-32 with Sun method
     *
     * @param ba byte array to compute CRC on
     *
     * @return 32-bit CRC, signed
     */
    private static int sunCRC32( byte[] ba ) {
        // create a new CRC-calculating object
        final CRC32 crc = new CRC32();
        crc.update( ba );
        // crc.update( int ) processes only the low order 8-bits. It actually expects an unsigned byte.
        return ( int ) (crc.hashCode());
    }

    private static long calculateCRC32(byte[] ba) {
        Checksum checksum = new CRC32();

        checksum.update(ba, 0, ba.length);
        return checksum.getValue();
    }



    public static void main(String[] args) {
        System.out.println("Hello");
         /*
          *       Wczytujemy źródłowy plik tekstowy 'Z'
          */
        BufferedReader in = null;
        BufferedWriter out = null;
        try {
            in = new BufferedReader(new FileReader("./Z.txt"));
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        }

        try {
            out = new BufferedWriter(new FileWriter("./W.txt"));
        } catch (IOException e) {
            e.printStackTrace();
        }

        String line, frame;
        try {
            while ((line = in.readLine()) != null) {
                /*
                 *  Dzielimy tekst na porcje po 32 bity
                 */
                int m = 32;

                for (int n = 0; n < line.length(); n+=m) {
                    if( n+m > line.length()) {
                        m = line.length() - n;
                    }
                    frame = line.substring(n,n+m);
                    /*
                    Obliczamy pole kontrolne CRC
                     */
                    final byte[] ba = frame.getBytes(UTF8Charset);
                    long bb = calculateCRC32(ba);
                    String crc = Long.toBinaryString(bb);
                    frame = biteStuffing(frame);

                    frame = "01111110" + frame + crc + "01111110";
                    System.out.println(frame +" " +crc.length());
                    out.write(frame);
                    out.write("\n");
                }

            }
            out.close();
        } catch (IOException e) {
            e.printStackTrace();
        }

//        try {
//            while ((line = in.readLine()) != null)   {
//                /*
//                 *      Dzielimy tekst na ramki, 8 bitów każda  (S tekst T)
//                 */
//                //public String substring(int startIndex,int endIndex)
//                int m = 8;
//                for (int n = 0; n < line.length(); n+=8) {
//                    frame = "S"+line.substring(n, m);
//                    /*
//                     *      Dla każdej ramki obliczamy pole kontrolne CRC i wstawiamy do ramki
//                     */
//                    final byte[] ba = frame.getBytes( UTF8Charset );
//                    //System.out.println( sunCRC32( ba ) );
//                    int bb = sunCRC32( ba );
//                    //System.out.println(Integer.toBinaryString(bb));
//                    String bc = Integer.toBinaryString(bb); //suma kontrolna wyrażona binarnie
//                    frame = frame + "E" + bc + "T"; //miedzy Escape a Termination stoi CRC
//                    /*
//                     *      Ramki zapisujemy kolejno do pliku tekstowego 'W'
//                     */
//                    System.out.println(frame);
//                    out.write(frame);
//                    m += 8;
//                    out.write('\n');
//                }
//                out.close();
//            }
//        } catch (IOException e) {
//            e.printStackTrace();
//        }

//        try {
//            while ((line = in.readLine()) != null)   {
//                /*
//                 *      Dzielimy tekst na ramki, 8 bitów każda  (S tekst T)
//                 */
//                //public String substring(int startIndex,int endIndex)
//                int m = 8;
//                for (int n = 0; n < line.length(); n+=8) {
//                   // frame = Integer.toBinaryString('S') +line.substring(n, m);
//                    frame = "S" + line.substring(n, m);
//                    System.out.println(frame);
//                    /*
//                     *      Dla każdej ramki obliczamy pole kontrolne CRC i wstawiamy do ramki
//                     *      obliczamy razem z S czy bez niego ?? chyba razem
//                     */
//                    final byte[] ba = frame.getBytes( UTF8Charset );
//                    System.out.println( new String (ba, UTF8Charset) );
//                    //System.out.println( sunCRC32( ba ) );
//                    int bb = sunCRC32( ba );
//                    //System.out.println(Integer.toBinaryString(bb));
//                    String bc = Integer.toBinaryString(bb); //suma kontrolna wyrażona binarnie
//                    frame = frame + "E" + bc + "T"; //miedzy Escape a Termination stoi CRC
//                                        /*
//                                         *      Ramki zapisujemy kolejno do pliku tekstowego 'W'
//                                         */
//                    System.out.println(frame);
//                    out.write(frame);
//                    m += 8;
//                    out.write('\n');
//                }
//                out.close();
//            }
//        } catch (IOException e) {
//            e.printStackTrace();
//        }

    }

    private static String biteStuff(String data) {
        if (data.contains("11111")) {
            String[] d = data.split("11111");
            if (d.length > 0) {
                if (d.length > 2) {
                    for (int i = 1; i < d.length; i++) {
                        d[i] = "0" + d[i];
                    }
                }
                String ret = "";
                for (int i = 0; i < d.length; i++) {
                    if (i + 1 >= d.length) {
                        ret = ret + d[i];
                    } else {
                        ret = ret + d[i] + "11111";
                    }
                }
                return ret;
            } else return "kok";
        } else return data;
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
}
