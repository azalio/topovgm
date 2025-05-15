# TopoVGM — Быстрый старт (РУССКИЙ)

TopoVGM — это Kubernetes-оператор для управления группами томов LVM на узлах кластера.

## Требования

- go v1.22.0+
- docker 17.03+
- kubectl v1.30+
- Доступ к Kubernetes v1.30+ кластеру
- lsblk (util-linux 2.39.4+)
- lvm2 (2.03.11+, желательно 2.03.23) на каждом узле

### Установка зависимостей на узлах

**Debian/Ubuntu:**

```sh
sudo apt-get install -y lvm2 util-linux
```

**CentOS/Fedora/RHEL:**

```sh
sudo dnf install -y lvm2 util-linux
```

---

## Установка TopoVGM в Kubernetes

1. **Соберите и опубликуйте Docker-образ:**

   ```sh
   make docker-build docker-push IMG=<ваш-реестр>/topovgm:tag
   ```

   Замените `<ваш-реестр>/topovgm:tag` на адрес вашего Docker-реестра.

2. **Установите CRD (CustomResourceDefinition):**

   ```sh
   make install
   ```

3. **Задеплойте оператор с вашим образом:**

   ```sh
   make deploy IMG=<ваш-реестр>/topovgm:tag
   ```

4. **(Опционально) Примените пример CR для теста:**

   ```sh
   kubectl apply -k config/samples/
   ```

---

## Проверка установки

- Проверьте, что CRD установлены:

  ```sh
  kubectl get crd | grep volumegroup
  ```

- Посмотрите логи оператора:

  ```sh
  kubectl -n <namespace> logs deployment/topovgm-controller-manager
  ```

  (Замените `<namespace>` на нужный, если не используется default)

---

## Удаление

1. Удалите примеры CR:

   ```sh
   kubectl delete -k config/samples/
   ```

2. Удалите CRD:

   ```sh
   make uninstall
   ```

3. Удалите оператор:

   ```sh
   make undeploy
   ```

---

## Дополнительно

- Для получения справки по make-целям:

  ```sh
  make help
  ```

- Подробнее о Kubebuilder: <https://book.kubebuilder.io/introduction.html>

---

## Лицензия

Проект лицензирован под Apache License 2.0.

## Важно

TopoVGM использует библиотеку **lvm2go** для управления LVM (создание, удаление, изменение Volume Group и т.д.), не вызывая напрямую внешние команды LVM. Однако для получения информации о блочных устройствах используется внешняя команда **lsblk** (через util-linux).

**lvm2** и **util-linux (lsblk)** должны быть установлены только на тех узлах, где будет работать оператор, чтобы lvm2go и lsblk могли выполнять свои функции.
