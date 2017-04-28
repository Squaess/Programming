package main;

import java.io.BufferedReader;
import java.util.Random;
import java.io.FileReader;
import java.io.IOException;

/**
 * Created by aedd on 4/27/17
 */
public class Main {
    private static final String FILE_PATH = "./data.txt";
    public static void main(String[] args) {
        BinaryTree<Integer> binaryTree = new BinaryTree<>();

        try {
            BufferedReader input = new BufferedReader(new FileReader(FILE_PATH));
            String line = input.readLine();
            int n = Integer.parseInt(line);
            for(int i = 0; i < n; i++) {
                line = input.readLine();
                String[] parts = line.split(" ");
                switch ( parts[0] ) {
                    case "insert"   :  // System.out.println("Insert "+ parts[1]);
                                        binaryTree.insert(Integer.parseInt(parts[1]),binaryTree.root);
                                        break;
                    case "delete"   :  // System.out.println("Delete " + parts[1]);
                                        Node<Integer> tmp = binaryTree.search(Integer.parseInt(parts[1]), binaryTree.root);
                                        binaryTree.delete(tmp);
                                        break;
                    case "min"      :  // System.out.println("min");
                                        binaryTree.min();
                                        break;
                    case "max"      :  // System.out.println("max");
                                        binaryTree.max();
                                        break;
                    case "inorder"  :  // System.out.println("in_order");
                                        binaryTree.draw(binaryTree.root);
                                        System.out.print("\n");
                                        break;
                    case "find"     :  // System.out.println("find " + parts[1]);
                                        if(binaryTree.search(Integer.parseInt(parts[1]), binaryTree.root) != null) System.out.println("1");
                                        else System.out.println("0");
                                        break;
                }
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
//        experiment();
    }

    private static void experiment() {
        int maximum_size_of_data = 1000000;
        int number_of_test_for_search = 100;

        for(int j = 0; j < maximum_size_of_data; j+= 100) {

            BinaryTree<Integer> bt = new BinaryTree<>();
            long min = Integer.MAX_VALUE;
            long max = 0;
            long avg = 0;

            /*
              Ilosc wezłów dla naszego drzewa
             */
            Random r = new Random();
            int value = r.nextInt();

            /*
              Wypełnienie drzewa losowymi liczbami
             */
            for (int i = 0; i < j; i++) {
                bt.insert(value-i, bt.root);
    //            bt.insert(r.nextInt(), bt.root);
            }

            for(int k = 0; k < number_of_test_for_search; k++) {
    //            bt.search(r.nextInt(), bt.root);
                bt.search(value-j+1, bt.root);
                if(bt.number_of_compare > max) max = bt.number_of_compare;
                if(bt.number_of_compare < min) min = bt.number_of_compare;
                avg += bt.number_of_compare;
                bt.number_of_compare = 0;
            }
            avg /= number_of_test_for_search;
            System.out.println(j+ " " +min+ " " +avg+ " " +max);

        }
    }
}
