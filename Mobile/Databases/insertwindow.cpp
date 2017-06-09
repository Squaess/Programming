#include "insertwindow.h"
#include "ui_insertwindow.h"
#include "mainwindow.h"
#include <QDebug>

InsertWindow::InsertWindow(QWidget *parent) :
    QMainWindow(parent),
    ui(new Ui::InsertWindow)
{
    ui->setupUi(this);
}

InsertWindow::~InsertWindow()
{
    delete ui;
}

void InsertWindow::on_addB_clicked()
{
  QString data = "";
  data += ui->nameF->text();
  data += " ";
  data += ui->surnameF->text();
  qDebug() << data;
  emit plsAddStud(data);
}
