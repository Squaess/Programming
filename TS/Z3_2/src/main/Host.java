package main;

import java.util.*;
/**
 * Created by Bartosz on 16.05.2017
 */
public class Host {

    private int signalDuration;
    private int cable_position;
    private char message;
    private int transsmisionDuration;
    private boolean isSending;
    private boolean error;

    private int errors_count;
    private int counter;

    private double probability;


    Host(int signalDuration, int cable_position, char message, double probability) {
        this.cable_position = cable_position;
        this.signalDuration = signalDuration;
        this.message = message;
        transsmisionDuration = 0;
        isSending = false;
        error = false;
        errors_count = 0;
        counter = 0;
        this.probability = probability;
    }

    int getCable_position() {
        return cable_position;
    }

    char getMessage() {
        return message;
    }

    void start(){
        isSending = true;
        this.transsmisionDuration = Ethernet.getLength()*2;
    }

    Cable_unit createUnit(Direction direction){
        if(transsmisionDuration == Ethernet.getLength()*2){
            if(error){
                System.out.print("\tRestart "+message);
            } else
            System.out.print("\t"+cable_position+" Starting transmission");
        }
        transsmisionDuration--;
        return new Cable_unit(this,cable_position,signalDuration,direction);
    }

    boolean wantToTransmit(){
        if(transsmisionDuration == 0){
            if(isSending){
                System.out.print("\t Correct"+message);
            }
            isSending = false;
            error = false;
        }
        return isSending;
    }

    void error(){
        signalDuration = Ethernet.getLength()*2;
        error = true;

        if(errors_count == 0){
            errors_count = 2;
            counter = 2;
        } else if(errors_count<1024){
            errors_count = errors_count*2;
            int random = new Random().nextInt(errors_count);
            counter = random;
        }
        System.out.print("\tStopped "+message+" ("+counter+")");
    }

    void waitForTransmission(){
        if(counter > 0) counter--;
    }

    boolean isWaiting(){
        if(counter>0){
            isSending = true;
            return true;
        }
        return false;
    }

    double getProbability() {
        return probability;
    }
}
