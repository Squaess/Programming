#ifndef MAINWINDOW_H
#define MAINWINDOW_H

#include <QMainWindow>
#include <QPushButton>
#include <QLabel>
#include <QVBoxLayout>

namespace Ui {
    class MainWindow;
}

class MainWindow : public QMainWindow
{
    Q_OBJECT
public:
    explicit MainWindow(QWidget *parent = 0);
    ~MainWindow();
private slots:
    void handleButton();
private:
    QPushButton *m_button;
    QPushButton *clearB;
    QPushButton *nextB;
    QPushButton *checkB;
    QLabel *correct;
    QLabel *failed;
};

#endif // MAINWINDOW_H
