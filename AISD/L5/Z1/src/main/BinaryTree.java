package main;

/**
 * Created by aedd on 5/10/17
 */

public class BinaryTree <T extends Comparable<T>>{

    private long number_of_compare = 0;
    Node<T> root;

    /**
     * konstruktor drzewa
     */
    BinaryTree()
    {
        root = null;
    }

    /**
     * wyszykiwanie wartosci
     * @param value wartosc
     * @param up gdzie zaczac(root)
     * @return wyszukany node, brak - null
     */
    Node<T> search(T value, Node<T> up) {
        if(up!=null){
            number_of_compare++;
            if(up.getValue().compareTo(value)==0) {

                return up;
            } else {
                number_of_compare++;
                if (up.getValue().compareTo(value) >= 0) {

                    return search(value, up.getLeft());
                } else
                    return search(value, up.getRight());
            }
        }
        return null;
    }
    /**
     * dodawanie wezla
     * @param value wartosc
     * @param up gdzie zaczassz przeszukiwanie(root)
     */
    void insert(T value, Node<T> up) {
        if(root == null) root = new Node<>(value);
        else{
            if( up.getValue().compareTo(value) >= 0){
                if(up.getLeft() != null)
                    insert(value, up.getLeft());
                else
                    up.setLeft(new Node<>(value, up));
            }
            else if(up.getValue().compareTo(value) < 0){
                if(up.getRight() != null)
                    insert(value, up.getRight());
                else
                    up.setRight(new Node<>(value, up));
            }
        }
    }
    /**
     * usowanie wezla
     * @param up node ktory usunac
     */
    void delete(Node<T> up) {
        if(up!=null){
            if(up.getLeft()==null && up.getRight()==null){
                if(up.getParent()==null) root=null;  //up is root
                else if(up.getParent().getLeft()==up) up.getParent().setLeft(null);
                else if(up.getParent().getRight()==up)up.getParent().setRight(null);
            }
            else if(up.getLeft()==null || up.getRight()==null){
                if(up.getLeft()==null){
                    if(up.getParent()==null){  //up is root
                        root = up.getRight();
                        root.setParent(null);
                    }
                    else if(up.getParent().getRight()==up){
                        up.getParent().setRight(up.getRight());
                        up.getRight().setParent(up.getParent());
                    }
                    else if(up.getParent().getLeft()==up)
                    {
                        up.getParent().setLeft(up.getRight());
                        up.getRight().setParent(up.getParent());
                    }
                }
                else if(up.getRight()==null){
                    if(up.getParent()==null){//up is root
                        root = up.getLeft();
                        root.setParent(null);
                    }
                    else if(up.getParent().getRight()==up){
                        up.getParent().setRight(up.getLeft());
                        up.getLeft().setParent(up.getParent());
                    }
                    else if(up.getParent().getLeft()==up)
                    {
                        up.getParent().setLeft(up.getLeft());
                        up.getLeft().setParent(up.getParent());
                    }
                }
            }
            else{
                Node<T> help = minRight(up.getRight());
                up.setValue(help.getValue());
                delete(help);
            }
        }
    }
    /**
     * wyszukiwanie najmniejszego nastepnika
     * @param up gdzie zaczac(root)
     * @return najwiekszey node z lewego drzewa up
     */
    private Node<T> minRight(Node<T> up){
        if(up.getLeft()==null) return up;
        else return minRight(up.getLeft());
    }
    /**
     * rysowanie
     * @param up gdzie zaczac(root)
     *
     */
//    public void draw(Node<T> up, StringBuilder s) {
//        if(up!=null){
//            if(up.getLeft()!=null) draw(up.getLeft(),s);
//            //System.out.print(up.getValue()+" ");
//            s.append(" "+up.getValue());
//            if(up.getRight()!=null) draw(up.getRight(),s);
//        }
//    }
    void draw(Node<T> up) {
        if(up!=null) {
            if(up.getLeft()!=null) {
                draw(up.getLeft());
            }
            System.out.print(up.getValue()+" ");
            if(up.getRight() != null) draw(up.getRight());
        }
    }

    void min() {
        if(root == null) System.out.println("");
        else {
            Node n = root;
            while(n.getLeft() != null) {
                n = n.getLeft();
            }
            System.out.println(n.getValue());
        }
    }

    void max(){
        if(root == null) System.out.println("");
        else {
            Node n = root;
            while(n.getRight() != null) {
                n = n.getRight();
            }
            System.out.println(n.getValue());
        }
    }
}
