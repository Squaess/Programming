package main;

/**
 * Created by Bartosz on 16.05.2017
 */
class Cable_unit extends Thread {

    Cable_unit left = null;
    Cable_unit right = null;
    String value;
    String side;

    Cable_unit(String value) {
        this.value = value;
    }
    Cable_unit() {
        this.value = "-";
    }

    @Override
    public void run() {

    }
}
