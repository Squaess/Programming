package main;

/**
 * Created by aedd on 4/27/17
 */
class Node<T> {
    private Node<T> left;
    private Node<T> right;
    private Node<T> parent;
    private T value;

    /**
     * konstruktor wezla
     * @param value wartosc rozpoznawcza
     */
    Node(T value) {
        this.value = value;
        this.left = null;
        this.right = null;
        this.parent = null;
    }

    /**
     * konstruktor
     * @param value wartosc rozpoznawcza
     * @param parent rodzic
     */
    Node(T value, Node<T> parent) {
        this.value = value;
        this.left = null;
        this.right = null;
        this.parent = parent;
    }

    /**
     * zwraca lewego syna
     * @return lewy syn
     */
    Node<T> getLeft(){
        return left;
    }
    /**
     * zwraca prawego syna
     * @return prawy syn
     */
    Node<T> getRight(){
        return right;
    }
    /**
     * zwraca rodzica
     * @return rodzic
     */
    Node<T> getParent(){
        return parent;
    }
    /**
     * zwraca wartosc rozpoznawcza
     * @return wartosc
     */
    T getValue(){
        return value;
    }
    /**
     * zmienia wartosc node
     * @param value wartosc
     */
    void setValue(T value){
        this.value = value;
    }
    /**
     * zmienia lewego syna
     * @param left lewy syn
     */
    void setLeft(Node<T> left){
        this.left = left;
    }
    /**
     * zmienia  prawego syna
     * @param  right prawy syn
     */
    void setRight(Node<T> right){
        this.right = right;
    }
    /**
     * zmienia rodzica
     * @param parent rodzic
     */
    void setParent(Node<T> parent){
        this.parent = parent;
    }
}
