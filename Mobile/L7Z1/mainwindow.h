#ifndef MAINWINDOW_H
#define MAINWINDOW_H

#include <QMainWindow>
#include "mydialog.h"

namespace Ui {
class MainWindow;
}

class MainWindow : public QMainWindow
{
    Q_OBJECT

public:
    explicit MainWindow(QWidget *parent = 0);
    ~MainWindow();

    void setName(const QString &name);
    QString name() const;

private slots:
    void on_nextButton_clicked();
    void on_checkButton_clicked();

    void on_showButton_clicked();

private:
    Ui::MainWindow *ui;
    MyDialog *mDialog;
    QString dictionary[21][2]{
        {"home","dom"},
        {"tree","drzewo"},
        {"screen","obraz"},
        {"accountant","ksiegowa"},
       {"actor", "aktor"},
       {"actress","aktorka"},
       {"air traffic controller","kontroler lotow"},
       {"architect","architekt"},
       {"artist","artysta"},
       {"attorney","adwokat"},
       {"banker","bankier"},
       {"bartender","barman"},
       {"barber","fryzjer"},
       {"bookkeeper","ksiegowy"},
       {"builder","budowniczy"},
       {"businessman","biznesmen"},
       {"businesswoman","biznesmenka"},
       {"businessperson","biznesmen"},
       {"butcher","rzeznik"},
       {"carpenter","ciesla"},
       {"cashier","kasjer"}
    };
};

#endif // MAINWINDOW_H
