package main;

import java.util.Random;

/**
 * Created by aedd on 5/10/17
 */

public class BstOS {

    public static void main (String[] args) {
//        BinaryTree<Integer> bst = new BinaryTree<>();
//        bst.insert(12, bst.root);
//        bst.insert(100, bst.root);
//        bst.insert(150, bst.root);
//        bst.insert(0, bst.root);
//        bst.insert(8, bst.root);
//        bst.insert(-2, bst.root);
//        bst.insert(50, bst.root);
//        bst.insert(17, bst.root);
//        bst.insert(18, bst.root);
//        bst.draw(bst.root);
//        System.out.print("\n");
//        System.out.println(bst.OS_Select(bst.root, 9));
//        System.out.println(bst.OS_Rank(bst.search(150, bst.root)));
//        bst.delete(bst.search(12, bst.root));
//        bst.delete(bst.search(50, bst.root));
//
//        bst.insert(50, bst.root);
//        bst.insert(12, bst.root);
//
//        System.out.println(bst.OS_Select(bst.root, 9));
//        System.out.println(bst.OS_Rank(bst.search(150, bst.root)));
//
//        BinaryTree<Integer> b = new BinaryTree<>();
//        b.insert(13, b.root);
//        b.insert(10, b.root);
//        b.insert(12, b.root);
//        b.draw(b.root);
//        System.out.print("\n");
//        b.delete(b.search(10,b.root));
//        b.draw(b.root);
        Random r = new Random();

        for (int n = 1; n < 10000; n += 100) {
            int compare = 0;
            for (int i = 0; i < 100; i++) {
                BinaryTree<Integer> bst = new BinaryTree<>();
                for (int j = 0; j<n; j++) {
                    bst.insert(r.nextInt(), bst.root);
                }
                int cmp = 0;
                for (int k = 0; k<100; k++) {
                    bst.OS_Select(bst.root, r.nextInt(n)+1);
                    cmp += bst.number_of_compare;
                }
                cmp /= 100;
                compare += cmp;

            }
            compare /= 100;
            System.out.println(n+" " + compare);
        }
    }
}
