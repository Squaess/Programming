package main;

import java.io.*;
import java.nio.charset.Charset;
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

        String s = "S";
        byte[] b = s.getBytes(UTF8Charset);
        System.out.println(b[0]);

        String s1 = "01010011";

        int i = Integer.parseInt(s1,2);

        System.out.println(s1);
        byte[] ba =  s1.getBytes(UTF8Charset) ;
        System.out.println(ba[0]);
        long n = calculateCRC32(b);


        System.out.println(n);
        n = calculateCRC32(b);
        System.out.println(n);

    }
}
