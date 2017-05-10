package main;

/**
 * Created by aedd on 5/10/17
 */

public class BstOS {

    public static void main (String[] args) {
        BinaryTree<Integer> bst = new BinaryTree<>();
        bst.insert(3, bst.root);
        bst.insert(5, bst.root);
        bst.draw(bst.root);
        System.out.print("\n");
        bst.max();
        bst.min();
    }
}
