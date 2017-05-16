package main;

/**
 * Created by Bartosz on 16.05.2017
 */
public class Ethernet {

    private static Host[] hosts;
    private static Medium medium;

    public static void main(String[] args) {

        medium = new Medium();
        medium.start();
        hosts = new Host[3];
        for (int i = 0; i<hosts.length; i++) {
            hosts[i] = new Host(i, medium.getSlot(), medium);
            hosts[i].start();
        }
    }
}
