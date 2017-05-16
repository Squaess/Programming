package main;

import java.util.*;
/**
 * Created by Bartosz on 16.05.2017
 */
public class Host extends Thread {
    int id;
    int cable_position;
    Medium medium;

    Host(int id, int cable_position, Medium m) {
        super();
        this.id = id;
        this.cable_position = cable_position;
        this.medium = m;
    }

    @Override
    public void run(){
        while(true) {

            Random r = new Random();
            int w8 = r.nextInt(6) + 1;
            try {
                Thread.sleep(w8 + 2000);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
            medium.isSending[id] = true;
        }
    }

}
