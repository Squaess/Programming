package main;

import java.util.ArrayList;
import java.util.List;
import java.util.Random;

/**
 * Created by aedd on 5/10/17
 */
public class Network {
    private static char empty = '-';

    private static int length = 70;

    private static List<Host> hosts = new ArrayList<>();
    private static List<Signal> signals = new ArrayList<>();

    private static char[] cable = null;

    /**
     *
     * @return ilosc komorek w kablu
     */
    static int getLength()
    {
        return length;
    }

    /**
     *
     * @param position pozycja hosta na kablu
     * @param message   wiadomosc jaka host przekazuje
     * @param signalDuration    czas trwania sygnalu
     */
    private static void addHost(int position, char message, int signalDuration)
    {
        Host host = new Host(position, message, signalDuration);
        hosts.add(host);

        System.out.println("Dodano użytkownika " + message + " na pozycji " + position + ".");
    }

    /**
     * Funkcja inicjujaca hosty oraz kabel
     */
    private static void initialize()
    {
        addHost(0, 'A', length+1);
        addHost(10, 'B', length+1);
        addHost(19, 'C', length+1);

        if(cable == null)
        {
            cable = new char[length];
            for(int i=0; i<cable.length; i++) cable[i] = empty;
        }
    }

    /**
     * funkcja realizujaca krok na kablu
     */
    private static void loop()
    {
        // Likwidacja sygnalow ktore juz skonczyly
        for(int i=0; i<signals.size(); i++)
        {
            if(signals.get(i).getSignalDuration() <= 0)
                signals.remove(i);
        }

        // Tablica przechowujaca nalozone sygnaly
        List<String> array = new ArrayList<>();
        for(int j=0; j<length; j++) array.add("");

        char error = '*';
        for (Signal signal : signals) {
            int position = signal.getPosition();
            String message = Character.toString(error);
            if (signal.getMaster() != null)
                message = Character.toString(signal.getMaster().getMessage());

            if (!array.get(position).contains(message)) {
                String current = array.get(position);
                String changed = current + message;
                array.set(position, changed);
            }
        }

        // Nanoszenie poprawek na kabel i drukowanie
        for(int i=0; i<cable.length; i++)
        {
            if(array.get(i).length() == 0)
                cable[i] = empty;
            else if(array.get(i).length() == 1)
                cable[i] = array.get(i).charAt(0);
            else
                cable[i] = error;

            System.out.print(cable[i]);
        }

        //Sprawdzanie czy dojdzie do kolizji sygnałów. Jeżeli tak to tworzy się zakłócenie
        for(int i=0; i<signals.size(); i++)
        {
            for (Signal signal : signals) {
                Signal p = signals.get(i);

                //PQ lub P_Q
                if (p.getPosition() + 1 == signal.getPosition() || p.getPosition() + 2 == signal.getPosition()) {
                    if (p.getDirection() && !signal.getDirection()) {
                        if (p.getMaster() != signal.getMaster()) {
                            signals.get(i).lose();
                            signal.lose();
                        }
                    }
                }
            }
        }

        // Losowanie hostów do wysłania wiadomości
        for (Host host1 : hosts) {
            double probability = new Random().nextDouble();
            double p = 0.01;
            if (probability < p) {
                host1.start();
            }
        }

        //Hosty - rozpatrzenie przypadków powiadomienie o błedach
        for (Host host : hosts) {
            int position = host.getPosition();
            char message = host.getMessage();

            if (host.isWaiting()) {
                host.waitForTransmission();
            } else {
                if (host.wantToTransmit()) {
                    if (cable[position] == empty || cable[position] == message) {
                        signals.add(host.createSignal(true));
                        signals.add(host.createSignal(false));
                    } else {
                        host.error();
                    }
                }
            }
        }

        //Przesuwanie sygnałów
        for (Signal signal : signals) {
            signal.move();
        }
    }

    public static void main(String[] args)
    {
        initialize();
        for(int i=0; i<1000; i++)
        {
            System.out.print("\n" + i + "\t");
            loop();
        }
    }
}
