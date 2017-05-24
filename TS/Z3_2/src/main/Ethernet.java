package main;

import java.util.ArrayList;
import java.util.List;
import java.util.Random;

/**
 * Created by Bartosz on 16.05.2017
 */
public class Ethernet {

    private static char clear = '-';
    private static char error = '*';

    private static int length = 150;

    private static List<Host> hosts = new ArrayList();
    private static List<Cable_unit> signals = new ArrayList<>();

    private static char[] cable = null;

    /**
     *
     * @param signalDuration ilosc migniec sygnalu
     * @param position miejsce w tablicy na ktorym ma nadawac
     * @param sign wiadomosc ktora bedzie propagowana w kablu
     */
    private static void addHost(int signalDuration, int position, char sign, double probability ) {
        Host host = new Host(signalDuration,position,sign,probability);
        hosts.add(host);

        System.out.println("Added host "+position+" message: "+sign);
    }

    private static void tick() {

        for (int i = 0; i < signals.size(); i++){
            if(signals.get(i).getSignalDuration() <= 0) signals.remove(i);
        }

        List<String> arr = new ArrayList<>();
        for (int i =0; i<length; i++) arr.add("");

        for (int i = 0; i < signals.size(); i++){
            Cable_unit u = signals.get(i);
            String message = Character.toString(error);
            if(u.getMaster() != null) message = Character.toString(signals.get(i).getMaster().getMessage());

            if(!arr.get(u.getPosition()).contains(message)){
                String tmp = arr.get(u.getPosition());
                String changed = tmp+message;
                arr.set(u.getPosition(), changed);
            }
        }

        for(int i = 0; i < cable.length; i++){
            if(arr.get(i).length() == 0) cable[i] = clear;
            else if (arr.get(i).length() == 1) cable[i] = arr.get(i).charAt(0);
            else cable[i] = error;

            System.out.print(cable[i]);
        }

        for(int i = 0; i<signals.size(); i++){
            for (int j = 0; j<signals.size(); j++){
                Cable_unit first = signals.get(i);
                Cable_unit second = signals.get(j);

                if(first.getPosition()+1 == second.getPosition() || first.getPosition()+2 == second.getPosition()){
                    if(first.getDirection()== Direction.RIGHT && second.getDirection()==Direction.LEFT){
                        if(first.getMaster() != second.getMaster()){
                            signals.get(i).lost();
                            signals.get(j).lost();
                        }
                    }
                }
            }
        }

        for(int i = 0; i < hosts.size(); i++){
            double p = new Random().nextDouble();
            if(p<hosts.get(i).getProbability()){
                hosts.get(i).start();
            }
        }

        // wysylanie wiadomosci
        for(int i = 0; i < hosts.size(); i++){
            Host h = hosts.get(i);
            if(h.isWaiting()){
                h.waitForTransmission();
            } else {
                if(h.wantToTransmit()){
                    if(cable[h.getCable_position()] == clear || cable[h.getCable_position()] == h.getMessage()){
                        signals.add(hosts.get(i).createUnit(Direction.RIGHT));
                        signals.add(hosts.get(i).createUnit(Direction.LEFT));
                    } else {
                        hosts.get(i).error();
                    }
                }
            }
        }

        for (int i = 0; i<signals.size(); i++){
            signals.get(i).tick();
        }
    }

    public static void main(String[] args) {
        initialize();
        for (int i = 0; i<1000; i++){
            System.out.print("\n");
            tick();
        }

    }

    private static void initialize() {
        addHost(length,0,'A', 0.01);
        addHost(length,length/2,'B', 0.01);
        addHost(length,length-1,'C',0.01);

        if(cable == null) {
            cable = new char[length];
            for (int i = 0; i < cable.length; i++) cable[i] = clear;
        }
    }

    public static int getLength() {
        return length;
    }
}
