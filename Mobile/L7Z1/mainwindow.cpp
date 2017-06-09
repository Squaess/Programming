#include "mainwindow.h"
#include "ui_mainwindow.h"
#include "mydialog.h"

MainWindow::MainWindow(QWidget *parent) :
    QMainWindow(parent),
    ui(new Ui::MainWindow)
{
    ui->setupUi(this);
    srand(time(NULL));
    connect(ui->checkButton, SIGNAL(clicked(bool)), this, SLOT(on_nextButton_clicked()));
 //   connect(ui->wordL, SIGNAL(paste(QString)), ui->nextButton, SLOT(on_nextButton_clicked()));
}

MainWindow::~MainWindow()
{
    delete ui;
}

void MainWindow::setName(const QString &name)
{
 ui->lineEdit->setText(name);
}

QString MainWindow::name() const
{
 return ui->lineEdit->text();
}

void MainWindow::on_nextButton_clicked()
{

    int i = rand() % 21;
    QString w = dictionary[i][0];
    ui->wordL->setText(w);
    ui->lineEdit->setText("");
}

void MainWindow::on_checkButton_clicked()
{
    QString w = ui->lineEdit->text();
    QString c = ui->wordL->text();
    int i = 0;
    for( i ; i < 21; i++){
        if(dictionary[i][0] == c) break;
    }
    if(dictionary[i][1]==w){
        int v = ui->correctL->text().split(" ")[1].toInt();
        v+=1;
        ui->correctL->setText("Correct: "+QString::number(v));
    } else {
        int v = ui->failedL->text().split(" ")[1].toInt();
        v+=1;
        ui->failedL->setText("Failed: "+QString::number(v));
    }
}

void MainWindow::on_showButton_clicked()
{
    QString w = ui->wordL->text();
    for(int i = 0; i<21; i++){
        if(dictionary[i][0] == w) {
            ui->lineEdit->setText(dictionary[i][1]);
        }
    }
}
